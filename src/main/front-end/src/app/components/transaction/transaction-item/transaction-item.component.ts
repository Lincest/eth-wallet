import {Component, OnInit} from '@angular/core';
import {ActivatedRoute} from "@angular/router";
import {TransactionService} from "../../../services/transaction.service";
import {Code} from "../../../models/resp";
import {MsgService} from "../../../services/msg.service";
import {TransactionResp} from "../../../models/transaction";

@Component({
  selector: 'app-transaction-item',
  templateUrl: './transaction-item.component.html',
  providers: [MsgService],
  styles: [    `
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

  constructor(
    private route: ActivatedRoute,
    private transactionService: TransactionService,
    private msgService: MsgService,
  ) {
  }

  ngOnInit(): void {
    this.loading = true;
    this.route.params.subscribe(res => {
      this.transactionHash = res['hash']
      this.transactionService.getTransactionByHash(this.transactionHash).subscribe(res => {
        this.loading = false;
        if (res.code === Code.err) {
          this.msgService.addError(`获取交易失败, ${res.msg}`)
        } else {
          this.data = res.data
        }
      })
    })
  }

}
