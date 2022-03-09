import { Component, OnInit } from '@angular/core';
import {LoginService} from "../services/login.service";
import {MessageService} from 'primeng/api';
import {Router} from "@angular/router";
import {Code} from "../models/resp";

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss'],
  providers: [MessageService]
})
export class LoginComponent implements OnInit {

  password: string = "";
  name: string = "";

  constructor(
    private loginService: LoginService,
    private msgService: MessageService,
    private router: Router,

  ) { }

  ngOnInit(): void {
  }

  logIn() {
    this.loginService.logIn(this.name, this.password).subscribe(res => {
      const ok = res.code === Code.ok
      // console.log('login ok = ', ok);
      if (ok) {
        this.msgService.add({severity: 'success', summary: '登录成功'});
        setTimeout(
          () => this.router.navigate(['/home']).then(), 500
        )
      } else {
        this.msgService.add({severity: 'error', summary: '登录失败', detail: res.msg});
      }
    });
  }

  register() {
    this.loginService.register(this.name, this.password).subscribe(res => {
      const ok = res.code === Code.ok
      if (ok) {
        this.msgService.add({severity: 'success', summary: '注册成功'});
        this.loginService.logIn(this.name, this.password).subscribe(res1 => {
          if (res1.code === Code.ok) {
            this.router.navigate(['/mnemonic']).then()
          } else {
            this.msgService.add({severity: 'error', summary: '登录失败', detail: res1.msg});
          }
        })
      } else {
        this.msgService.add({severity: 'error', summary: '注册失败', detail: res.msg});
      }
    })
  }

}
