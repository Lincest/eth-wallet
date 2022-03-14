import {Component, OnInit} from '@angular/core';
import {Network} from "../../models/network";
import {NetworkService} from "../../services/network.service";
import {Code} from "../../models/resp";
import {MsgService} from "../../services/msg.service";
import * as lodash from 'lodash';

@Component({
  selector: 'app-network',
  templateUrl: './network.component.html',
  providers: [MsgService]
})
export class NetworkComponent implements OnInit {

  items: Network[] = [];
  newNet: Network = {name: "", url: "", chain_id: ""};
  modifyDialogVisible: boolean = false;

  constructor(
    private networkService: NetworkService,
    private msgService: MsgService,
  ) {
  }

  ngOnInit(): void {
    this.loadNetWorks()
  }

  // 加载所有网络
  loadNetWorks() {
    this.networkService.getNetWorks().subscribe(res => {
      if (res.code === Code.err) {
        this.msgService.addError('获取网络节点信息失败')
      } else {
        this.msgService.addSuccess('获取网络节点信息成功')
        this.items = res.data as Network[]
      }
    })
  }

  // 修改网络
  modify(item: Network) {
    this.newNet = lodash.cloneDeep(item);
    this.modifyDialogVisible = true;
  }

  // 删除网络
  delete(item: Network) {
    this.msgService.confirm("您确定要删除吗?", () => {
      this.networkService.deleteNetwork(item).subscribe(res => {
        if (res.code === Code.err) {
          this.msgService.addError(res.msg)
        } else {
          this.msgService.addSuccess("成功删除网络")
          this.loadNetWorks()
        }
      })
    })
  }

  // 更换网络
  change(item: Network) {
    this.networkService.setCurrentNetwork(item).subscribe(res => {
      if (res.code === Code.ok) {
        this.msgService.addSuccess("成功切换网络")
      } else {
        this.msgService.addError(res.msg)
      }
    })
  }

  // 保存
  saveModification() {
    this.networkService.addNetwork(this.newNet).subscribe(res => {
      if (res.code === Code.err) {
        this.msgService.addError(res.msg)
      } else {
        this.msgService.addSuccess("成功添加网络")
        this.loadNetWorks()
        this.modifyDialogVisible = false;
      }
    })
  }

  // 添加网络
  addNet() {
    this.newNet = {name: "", url: "", chain_id: ""};
    this.modifyDialogVisible = true;
  }
}
