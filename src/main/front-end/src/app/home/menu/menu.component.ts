import { Component, OnInit } from '@angular/core';
import {ConfirmationService, MessageService} from "primeng/api";

@Component({
  selector: 'app-menu',
  templateUrl: './menu.component.html',
  providers: [MessageService, ConfirmationService]
})
export class MenuComponent implements OnInit {

  model: any[] = [];
  constructor() { }

  ngOnInit(): void {
    this.model = [
      {
        label: 'Home',
        items:[
          // {label: '登录 / 注册',icon: 'pi pi-fw pi-check', routerLink: ['/login']},
          {label: '主页',icon: 'pi pi-fw pi-home', routerLink: ['/home']},
        ],
      },
      {
        label: 'Applications',
        items: [
          {label: '网络', icon: 'pi pi-fw pi-cloud', routerLink: ['/home/network']},
          {label: '账户', icon: 'pi pi-fw pi-users', routerLink: ['/home/account']},
          {label: '创建交易', icon: 'pi pi-fw pi-money-bill', routerLink: ['/home/transaction']},
          {label: '检索交易', icon: 'pi pi-fw pi-search', routerLink: ['/home/transaction-search']},
          {label: '交易历史', icon: 'pi pi-fw pi-calendar', routerLink: ['/home/transaction-history']}
        ]
      },
  ]
  }

  onKeydown(event: KeyboardEvent) {
    const nodeElement = (<HTMLDivElement> event.target);
    if (event.code === 'Enter' || event.code === 'Space') {
      nodeElement.click();
      event.preventDefault();
    }
  }
}
