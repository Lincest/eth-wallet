import { Component, OnInit } from '@angular/core';
import {HomeComponent} from "../home.component";

@Component({
  selector: 'app-topbar',
  templateUrl: './topbar.component.html',
})
export class TopbarComponent implements OnInit {

  constructor(public home: HomeComponent) { }

  ngOnInit(): void {
  }

}
