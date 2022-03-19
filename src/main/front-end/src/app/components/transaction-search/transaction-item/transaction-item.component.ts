import {Component, OnInit} from '@angular/core';
import {ActivatedRoute, Router} from "@angular/router";
import {TransactionService} from "../../../services/transaction.service";
import {Code} from "../../../models/resp";
import {MsgService} from "../../../services/msg.service";
import {TransactionResp} from "../../../models/transaction";
import Web3 from "web3";

@Component({
  selector: 'app-transaction-item',
  templateUrl: './transaction-item.component.html',
  providers: [MsgService],
  styles: [`
    .badge {
      border-radius: var(--border-radius);
      padding: .25em .5rem;
      text-transform: uppercase;
      font-weight: 700;
      font-size: 12px;
      letter-spacing: .3px;

      &.success {
        background: #C8E6C9;
        color: #256029;
      }

      &.fail {
        background: #FFCDD2;
        color: #C63737;
      }

      &.pending {
        background: #FEEDAF;
        color: #8A5340;
      }
    }
  `]
})
export class TransactionItemComponent implements OnInit {

  loading = false;
  transactionHash: string;
  data: TransactionResp;

  // 加速交易表单弹窗
  accDialogVisible = false;
  gasPrice: string;
  suggestGasPrice: string = "0";

  constructor(
    private route: ActivatedRoute,
    private transactionService: TransactionService,
    private msgService: MsgService,
    private router: Router
  ) {
  }

  ngOnInit(): void {
    this.loading = true;
    this.route.params.subscribe(res => {
      this.transactionHash = res['hash']
      this.transactionService.getTransactionByHash(this.transactionHash).subscribe(res => {
        this.loading = false;
        if (res.code === Code.err) {
          this.router.navigate(['/404']).then()
        } else {
          this.data = res.data
        }
      })
    })
  }

  // 获取建议gas price
  getSuggestGasPrice(transaction: TransactionResp) {
    let web3 = new Web3(transaction.network)
    web3.eth.getGasPrice().then(
      res => {
        this.suggestGasPrice = Web3.utils.fromWei(res, 'gwei');
        this.gasPrice = Web3.utils.fromWei(res, 'gwei');
      }
    )
  }

  accelerateTransaction() {
    this.accDialogVisible = true;
    this.getSuggestGasPrice(this.data);
  }

  saveAccelerate() {
    const gasPriceWei = Web3.utils.toWei(this.gasPrice, 'gwei');
    console.log("new gas price in wei: ", gasPriceWei)
    this.transactionService.accelerateTransaction(gasPriceWei, this.data).subscribe(res => {
        if (res.code === Code.err) {
          this.msgService.addError(res.msg)
        } else {
          this.msgService.addSuccess()
          this.msgService.confirm("交易成功, 交易哈希值已经改变, 点击将跳转到新交易页", () => {
            this.accDialogVisible = false;
            this.router.navigate([`/home/transaction-search/${res.data}`]).then()
          })
        }
      }
    )
  }
}
