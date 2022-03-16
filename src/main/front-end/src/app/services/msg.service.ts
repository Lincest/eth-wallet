import { Injectable } from '@angular/core';
import {ConfirmationService, MessageService} from "primeng/api";

@Injectable({
  providedIn: 'root',
})
export class MsgService {

  constructor(
    public msgService: MessageService,
    public confirmService: ConfirmationService
  ) { }

  // 成功提示
  addSuccess(detail?: string) {
    const res = {key: 'home-toast', severity: 'success', summary: '成功', detail};
    this.msgService.add(res);
  }

  // 错误提示
  addError(detail?: string) {
    const res = {key: 'home-toast', severity: 'error', summary: '失败', detail};
    this.msgService.add(res);
  }

  // 确认框
  confirm(msg: string, accFunc: Function) {
    this.confirmService.confirm({
      message: msg,
      accept: () => accFunc()
    })
  }
}
