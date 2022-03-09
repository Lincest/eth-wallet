import { Component, OnInit } from '@angular/core';
import {LoginService} from "../services/login.service";
import {MessageService} from 'primeng/api';
import {Router} from "@angular/router";

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
    this.loginService.logIn(this.name, this.password).subscribe(ok => {
      // console.log('login ok = ', ok);
      if (ok) {
        this.msgService.add({severity: 'success', summary: '登录成功'});
        setTimeout(
          () => this.router.navigate(['/home']).then(), 500
        )
      } else {
        this.msgService.add({severity: 'error', summary: '登录失败'});
      }
    });
  }

}
