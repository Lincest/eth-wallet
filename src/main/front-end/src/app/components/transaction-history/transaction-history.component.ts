import { Component, OnInit } from '@angular/core';
import {TransactionResp} from "../../models/transaction";
import {TransactionService} from "../../services/transaction.service";
import {Code} from "../../models/resp";

@Component({
  selector: 'app-transaction-history',
  templateUrl: './transaction-history.component.html',
})
export class TransactionHistoryComponent implements OnInit {

  data: TransactionResp[];
  loading = false;
  totalRecords: number = 100;

  constructor(
    private transactionService: TransactionService
  ) { }

  ngOnInit(): void {
  }

  loadData(event: any) {
    this.loading = true;
    console.log("lazy event = ", event)
    const page_size = event.rows;
    const page = event.first / event.rows + 1;
    this.transactionService.getTransactionCount().subscribe(ans => {
      if (ans.code === Code.ok) {
        this.totalRecords = ans.data
        this.transactionService.getLatestTransaction({page, page_size}).subscribe(res => {
          this.loading = false;
          if (res.code === Code.ok) {
            this.data = res.data
            console.log(this.data)
          }
        })
      }
    })
  }
}
