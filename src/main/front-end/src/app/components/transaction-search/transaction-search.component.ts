import {Component, OnInit} from '@angular/core';
import {ActivatedRoute, Router} from "@angular/router";
import {Code} from "../../models/resp";
import {TransactionService} from "../../services/transaction.service";
import {MsgService} from "../../services/msg.service";

@Component({
  selector: 'app-transaction-search',
  templateUrl: './transaction-search.component.html',
  providers: [MsgService]
})
export class TransactionSearchComponent implements OnInit {

  loading = false;
  transactionHash: string;

  constructor(
    private route: ActivatedRoute,
    private router: Router,
    private transactionService: TransactionService,
    private msgService: MsgService,
  ) {
  }

  ngOnInit(): void {
    this.route.params.subscribe(p => {
      console.log('p = ', p)
      this.transactionHash = p['hash']
    })
  }

  checkTransaction() {
    this.loading = true;
    this.transactionService.getTransactionByHash(this.transactionHash).subscribe(res => {
      this.loading = false;
      if (res.code === Code.ok) {
        this.msgService.addSuccess()
        this.router.navigate([this.transactionHash], {relativeTo: this.route}).then()
      } else {
        this.msgService.addError(res.msg)
      }
    })
  }

}
