import {Component, OnInit} from '@angular/core';
import {NetworkService} from "../../services/network.service";
import {Account, defaultAccount} from "../../models/account";
import {AccountService} from "../../services/account.service";
import {Code} from "../../models/resp";
import Web3 from 'web3';
import {TransactionReq} from "../../models/transaction";
import {TransactionService} from "../../services/transaction.service";
import {MsgService} from "../../services/msg.service";

@Component({
  selector: 'app-transaction',
  templateUrl: './transaction.component.html',
  providers: [MsgService],
})
export class TransactionComponent implements OnInit {

  loading = false;

  // 网络
  network: string;

  // finished?
  finishedTransaction = false;
  transactionHash = "";

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

  // Gas Limit
  gasLimit: string = "21000";

  constructor(
    private networkService: NetworkService,
    private accountService: AccountService,
    private transactionService: TransactionService,
    private msgService: MsgService,
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
    let filtered: Account[] = [];
    let query = event.query;

    for (let i = 0; i < this.accounts.length; i++) {
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
    let filtered: Account[] = [];
    let query = event.query;

    for (let i = 0; i < this.accounts.length; i++) {
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

  // 提交
  submitTransaction() {
    const req: TransactionReq = {
      from_address: this.selectedFromAccount.address,
      from_private_key_hex: this.selectedFromAccount.private_key_hex,
      to_address: this.selectedToAccount.address,
      value: Web3.utils.toWei(this.transferValue, 'ether'),
      gas_price: Web3.utils.toWei(this.gasPrice, 'gwei'),
      gas_limit: this.gasLimit
    }
    this.msgService.confirm(`
        确认提交交易?
       `, () => {
      this.loading = true;
      this.finishedTransaction = false;
      this.transactionService.createTransaction(req).subscribe(res => {
        this.loading = false;
        if (res.code === Code.err) {
          this.msgService.addError(res.msg)
        } else {
          this.msgService.addSuccess(`发起交易成功`)
          this.transactionHash = res.data
          this.finishedTransaction = true;
        }
      })
    })
  }
}
