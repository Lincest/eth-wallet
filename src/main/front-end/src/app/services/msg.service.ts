import { Injectable } from '@angular/core';
import {MessageService} from "primeng/api";

@Injectable({
  providedIn: 'root',
})
export class MsgService {

  constructor(
    public msgService: MessageService
  ) { }

  addSuccess(detail?: string) {
    const res = {key: 'home-toast', severity: 'success', summary: '成功', detail};
    this.msgService.add(res);
  }

  addError(detail?: string) {
    const res = {key: 'home-toast', severity: 'error', summary: '失败', detail};
    this.msgService.add(res);
  }
}
