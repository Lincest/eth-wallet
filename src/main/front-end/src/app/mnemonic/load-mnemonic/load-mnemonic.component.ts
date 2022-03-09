import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-load-mnemonic',
  templateUrl: './load-mnemonic.component.html',
})
export class LoadMnemonicComponent implements OnInit {

  constructor() { }

  mnemonics: string[] = [];
  ngOnInit(): void {
  }

}
