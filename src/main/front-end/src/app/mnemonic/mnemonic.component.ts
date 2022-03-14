import { Component, OnInit } from '@angular/core';
import {ConfirmationService, Message, MessageService} from "primeng/api";
import {ActivatedRoute, Router} from "@angular/router";
import {WalletService} from "../services/wallet.service";
import {LoginService} from "../services/login.service";
import {Code} from "../models/resp";
import {errorPop, successPop} from "../models/global";
@Component({
  selector: 'app-mnemonic',
  templateUrl: './mnemonic.component.html',
  styleUrls: ['./mnemonic.component.scss'],
  providers: [MessageService, ConfirmationService]
})
export class MnemonicComponent implements OnInit {

  mnemonics: string[] = []; // 助记词
  constructor(
    private router: Router,
    private activeRoute: ActivatedRoute,
    private wallet: WalletService,
    private confirm: ConfirmationService,
    private msgService: MessageService,
    private loginService: LoginService
) { }

  ngOnInit(): void {
  }

  // 建立新助记词
  newMnemonic() {
    this.mnemonics = this.wallet.generateMnemonic()
  }

  // 导入旧助记词
  loadMnemonic() {
    if (!this.wallet.validateMnemonic(this.mnemonics)) {
      this.msgService.add({severity: 'error', summary: '助记词不合法', detail: '助记词不合法, 请检查您的助记词'});
      return
    }
    this.confirm.confirm({
      message: `您是否确定提交您的助记词？<br>
                您的助记词是: <br><br>
                ${this.mnemonics.join(" ")} <br></br>
                请务必用安全的介质存储助记词！`,
      accept: () => {
        console.log('acc');
        console.log(this.mnemonics);
        // TODO: 服务器记录助记词md5
        this.loginService.updateMnemonic(this.mnemonics.join(' ')).subscribe(res => {
          if (res.code === Code.ok) {
            this.msgService.add(successPop);
            setTimeout(() => {
              this.router.navigate(['/home']).then();
            }, 500)
          } else {
            this.msgService.add(errorPop);
          }
        })
      }
    })
  }
}
