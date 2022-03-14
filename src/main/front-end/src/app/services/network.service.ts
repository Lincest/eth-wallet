import { Injectable } from '@angular/core';
import {Observable} from "rxjs";
import {Network} from "../models/network";
import {AUTH_URL, BASE_URL} from "../models/global";
import {Resp} from "../models/resp";
import {HttpClient} from "@angular/common/http";

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
}
