import {Component, OnInit} from '@angular/core';
import {LoginService} from "./services/login.service";

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent implements OnInit {
  title = 'front-end';

  constructor(
    private loginService: LoginService
  ) {
  }
  ngOnInit() {
    this.loginService.isLoggedIn().subscribe(x => console.log(`登陆与否: ${x}`));
  }
}
