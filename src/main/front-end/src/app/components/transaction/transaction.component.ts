import { Component, OnInit } from '@angular/core';
import {NetworkService} from "../../services/network.service";
import {Account, defaultAccount} from "../../models/account";
import {AccountService} from "../../services/account.service";
import {Code} from "../../models/resp";
import Web3 from 'web3';

@Component({
  selector: 'app-transaction',
  templateUrl: './transaction.component.html',
})
export class TransactionComponent implements OnInit {

  // 网络
  network: string;

  // 地址
  accounts: Account[] = [];
  // from
  filteredFromAccounts: Account[] = [];
  selectedFromAccount: Account;
  // to
  filteredToAccounts: Account[] = [];
  selectedToAccount: Account;

  // transfer value
  transferValue: string = "0";

  // Gas Price
  gasPrice: string = "0";
  suggestGasPrice: string = "0";

  constructor(
    private networkService: NetworkService,
    private accountService: AccountService
  ) {
    this.selectedFromAccount = {...defaultAccount}
    this.selectedToAccount = {...defaultAccount}
  }

  ngOnInit(): void {
    this.networkService.getCurrentNetwork().subscribe(
      res => {
        this.network = res.url;
        this.getSuggestGasPrice();
      }
    )
    this.accountService.getAllAccounts().subscribe(res => {
      if (res.code === Code.ok) {
        this.accounts = res.data
      }
    })
  }

  // for auto complete选择: from address
  filterFromAddress(event: any) {
    //in a real application, make a request to a remote url with the query and return filtered results, for demo we filter at client side
    let filtered : Account[] = [];
    let query = event.query;

    for(let i = 0; i < this.accounts.length; i++) {
      let account = this.accounts[i];
      if (account.address.toLowerCase().indexOf(query.toLowerCase()) == 0) {
        filtered.push(account);
      }
    }
    this.filteredFromAccounts = filtered;
  }

  // for auto complete选择: to address
  filterToAddress(event: any) {
    //in a real application, make a request to a remote url with the query and return filtered results, for demo we filter at client side
    let filtered : Account[] = [];
    let query = event.query;

    for(let i = 0; i < this.accounts.length; i++) {
      let account = this.accounts[i];
      if (account.address.toLowerCase().indexOf(query.toLowerCase()) == 0) {
        filtered.push(account);
      }
    }
    this.filteredToAccounts = filtered;
  }

  // 获取建议web3
  getSuggestGasPrice() {
    let web3 = new Web3(this.network)
    const suggestGasPrice = web3.eth.getGasPrice().then(
      res => {
        this.suggestGasPrice = Web3.utils.fromWei(res, 'gwei');
        this.gasPrice = Web3.utils.fromWei(res, 'gwei');
      }
    )
  }
}
