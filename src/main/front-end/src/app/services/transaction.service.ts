import { Injectable } from '@angular/core';
import {TransactionReq} from "../models/transaction";
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

  createTransaction(req: TransactionReq): Observable<Resp> {
    return this.http.post<Resp>(AUTH_URL + "/transaction", req)
  }

  getTransactionByHash(hash: string): Observable<Resp> {
    return this.http.get<Resp>(AUTH_URL + `/transaction/${hash}`)
  }
}
