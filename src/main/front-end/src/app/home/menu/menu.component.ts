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
          {label: '网络', icon: 'pi pi-fw pi-cloud', routerLink: ['/home/network']}
        ]
      }

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
