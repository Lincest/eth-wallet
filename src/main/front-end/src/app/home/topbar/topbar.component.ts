import {Component, OnInit} from '@angular/core';
import {HomeComponent} from "../home.component";
import {MenuItem, MessageService} from "primeng/api";
import {LoginService} from "../../services/login.service";
import {Code} from "../../models/resp";
import {Router} from "@angular/router";

@Component({
  selector: 'app-topbar',
  templateUrl: './topbar.component.html',
  providers: [MessageService]
})
export class TopbarComponent implements OnInit {

  items: MenuItem[] = [];
  username: string = "";

  constructor(
    public home: HomeComponent,
    private loginService: LoginService,
    private msgService: MessageService,
    private router: Router
  ) {
  }

  ngOnInit(): void {
    const store = localStorage.getItem("wallet-login")
    if (store) {
      this.username = JSON.parse(store).name;
    }
    this.items = [
      {
        label: `欢迎您: ${this.username}`,
        escape: false,
        icon: 'pi pi-user',
      },
      {
        label: `<span class="text-orange-500 font-medium">登出</span>`,
        icon: `pi pi-sign-out`,
        escape: false,
        command: () => {
          this.loginService.logOut().subscribe(resp => {
            if (resp.code !== Code.ok) {
              this.msgService.add({severity: 'error', summary: '登出失败', detail: resp.msg});
            } else {
              this.msgService.add({severity: 'success', summary: '登出成功', detail: resp.msg});
              setTimeout(() => {
                this.router.navigate(['/login']).then();
              }, 500);
            }
          })
        }
      }
    ]

  }

}
