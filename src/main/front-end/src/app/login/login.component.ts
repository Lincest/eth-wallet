import { Component, OnInit } from '@angular/core';
import {LoginService} from "../services/login.service";
import {MessageService} from 'primeng/api';

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
    private msgService: MessageService
  ) { }

  ngOnInit(): void {
  }

  logIn() {
    this.loginService.logIn(this.name, this.password).subscribe(ok => {
      // console.log('login ok = ', ok);
      if (ok) {
        this.msgService.add({severity: 'success', summary: '登录成功'});
      } else {
        this.msgService.add({severity: 'error', summary: '登录失败'});
      }
    });
  }

}
