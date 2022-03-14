import {AfterViewInit, Component, OnInit} from '@angular/core';
import {WalletService} from "../services/wallet.service";
import {ConfirmationService, MessageService} from "primeng/api";

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.scss'],
  providers: [MessageService, ConfirmationService]
})
export class HomeComponent implements AfterViewInit, OnInit {

  menuClick = true;
  menuActive = false;
  profileActive = false;

  constructor(
    private wallet: WalletService,
    private msgService: MessageService
  ) {
  }

  ngOnInit(): void {
  }

  ngAfterViewInit() {
    this.menuClick = false;
  }

  toggleMenu(event: Event) {
    this.menuClick = true;
    this.menuActive = !this.menuActive;
    event.preventDefault();
  }

  toggleProfile(event: Event) {
    this.profileActive = !this.profileActive;
    event.preventDefault();
  }

}
