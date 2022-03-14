import { Injectable } from '@angular/core';
import {Observable} from "rxjs";
import {Network} from "../models/network";
import {AUTH_URL, BASE_URL} from "../models/global";
import {Resp} from "../models/resp";
import {HttpClient} from "@angular/common/http";
import {resolvePtr} from "dns";

@Injectable({
  providedIn: 'root'
})
export class NetworkService {

  constructor(
    private http: HttpClient
  ) { }

  // get
  getNetWorks(): Observable<Resp> {
    return this.http.get<Resp>(AUTH_URL + "/network")
  }

  // test
  isValidUrlAndChainId(): Observable<Resp> {
    return this.http.get<Resp>(BASE_URL + "/test-rpc", )
  }

  // delete
  deleteNetwork(item: Network): Observable<Resp> {
    return this.http.delete<Resp>(AUTH_URL + "/network", {body: item})
  }

  // add
  addNetwork(item: Network): Observable<Resp> {
    return this.http.post<Resp>(AUTH_URL + "/network", item)
  }
}
