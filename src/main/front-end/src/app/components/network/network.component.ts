import {Component, OnInit} from '@angular/core';
import {Network} from "../../models/network";
import {NetworkService} from "../../services/network.service";
import {Code} from "../../models/resp";
import {MsgService} from "../../services/msg.service";

@Component({
  selector: 'app-network',
  templateUrl: './network.component.html',
  providers: [MsgService]
})
export class NetworkComponent implements OnInit {

  items: Network[] = [];
  constructor(
    private networkService: NetworkService,
    private msgService: MsgService
  ) { }

  ngOnInit(): void {
    this.networkService.getNetWorks().subscribe(res => {
      if (res.code === Code.err) {
        this.msgService.addError()
      } else {
        this.msgService.addSuccess()
        this.items = res.data as Network[]
      }
    })
  }

}
