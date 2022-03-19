import { Injectable } from '@angular/core';
import {TransactionReq, TransactionResp} from "../models/transaction";
import {Observable} from "rxjs";
import {Resp} from "../models/resp";
import {AUTH_URL} from "../models/global";
import {HttpClient} from "@angular/common/http";

@Injectable({
  providedIn: 'root'
})
export class TransactionService {

  constructor(
    private http: HttpClient
  ) { }

  // 创建交易
  createTransaction(req: TransactionReq): Observable<Resp> {
    return this.http.post<Resp>(AUTH_URL + "/transaction", req)
  }

  // 查询交易
  getTransactionByHash(hash: string): Observable<Resp> {
    return this.http.get<Resp>(AUTH_URL + `/transaction/${hash}`)
  }

  // 加速交易
  accelerateTransaction(gasPrice: string, transaction: TransactionResp): Observable<Resp> {
    return this.http.post<Resp>(AUTH_URL + `/transaction/${transaction.ID}`, {gas_price: gasPrice})
  }
}
