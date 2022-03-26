import {Injectable, NgZone} from '@angular/core';
import {Observable} from "rxjs";
import {BASE_URL} from "../models/global";

@Injectable({
  providedIn: 'root'
})
export class SseService {

  constructor(
    private _zone: NgZone
  ) {
  }

  getEventSource(url: string): EventSource {
    return new EventSource(url);
  }

  // SSE support
  getServerSentEvent(url: string) {
    return new Observable(obs => {
      const eventSource = this.getEventSource(url);

      eventSource.onmessage = event => {
        this._zone.run(() => obs.next(event))
      }

      eventSource.onerror = err => {
        this._zone.run(() => obs.error(err))
      }
    })
  }

  // test
  test() {
    this.getServerSentEvent(BASE_URL + "/sse-test").subscribe((x:any) => console.log("event: ", x?.data))
  }


}
