import { Injectable } from '@angular/core';
import {TransactionReq, TransactionResp} from "../models/transaction";
import {Observable} from "rxjs";
import {Resp} from "../models/resp";
import {AUTH_URL} from "../models/global";
import {HttpClient, HttpParams} from "@angular/common/http";
import {Page} from "../models/page";

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

  // 根据tx hash查询交易
  getTransactionByHash(hash: string): Observable<Resp> {
    return this.http.get<Resp>(AUTH_URL + `/transaction/${hash}`)
  }

  // 加速交易
  accelerateTransaction(gasPrice: string, transaction: TransactionResp): Observable<Resp> {
    return this.http.post<Resp>(AUTH_URL + `/transaction/${transaction.ID}`, {gas_price: gasPrice})
  }

  // 查询当前用户当前网络最新交易
  getLatestTransaction(page?: Page): Observable<Resp> {
    let params = new HttpParams();
    if (page) {
      params = params.append('page', page.page);
      params = params.append('page_size', page.page_size);
    }
    return this.http.get<Resp>(AUTH_URL + "/transaction/latest", {params})
  }

  // 交易数目
  getTransactionCount(): Observable<Resp> {
    return this.http.get<Resp>(AUTH_URL + "/transaction/count")
  }
}
