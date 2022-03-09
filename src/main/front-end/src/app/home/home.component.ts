import {Component, OnInit} from '@angular/core';
import {WalletService} from "../services/wallet.service";

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.scss']
})
export class HomeComponent implements OnInit {

  constructor(
    private wallet: WalletService
  ) {
  }

  ngOnInit(): void {
  }


}
