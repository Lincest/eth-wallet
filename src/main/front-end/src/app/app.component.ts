import {Component, OnInit} from '@angular/core';
import {LoginService} from "./services/login.service";
import {PrimeNGConfig} from "primeng/api";

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent implements OnInit {
  title = 'front-end';

  constructor(
    private loginService: LoginService,
    private primengConfig: PrimeNGConfig
  ) {
  }
  ngOnInit() {
    // ripple
    this.primengConfig.ripple = true;
    // normal configs
    document.documentElement.style.fontSize = '14px';
    // check login
    this.loginService.isLoggedIn().subscribe(x => console.log(`登陆与否: ${x}`));
  }
}
