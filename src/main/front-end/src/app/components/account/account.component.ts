import {Component, OnInit} from '@angular/core';
import {Account} from "../../models/account";
import {AccountService} from "../../services/account.service";
import {MsgService} from "../../services/msg.service";
import {Code} from "../../models/resp";
import Web3 from 'web3';

@Component({
  selector: 'app-account',
  templateUrl: './account.component.html',
  providers: [MsgService],
  styles: [
    `
      .balance-badge {
        border-radius: var(--border-radius);
        padding: .25em .5rem;
        text-transform: uppercase;
        font-weight: 700;
        font-size: 12px;
        letter-spacing: .3px;
        background: #C8E6C9;
        color: #256029;
      }
    `
  ]
})
export class AccountComponent implements OnInit {

  accounts: Account[] = [];
  loading = false;

  // keystore导入
  keystoreImportVisible = false;
  keystorePassPhrase = "";

  // 私钥导入
  privateKeyImportVisible = false;
  privateKey = ""

  // 总余额

  totalBalance = ""

  constructor(
    private accountService: AccountService,
    private msgService: MsgService
  ) {
  }

  ngOnInit(): void {
    this.load()
  }

  load() {
    this.loading = true;
    this.accountService.getAllAccounts().subscribe(resp => {
      this.loading = false;
      if (resp.code === Code.ok) {
        this.accounts = resp.data as Account[];
        this.getTotalBalance();
      } else {
        this.msgService.addError("获取账户信息失败")
      }
    })
  }

  // 获取账户总余额
  getTotalBalance() {
    let value = Web3.utils.toBN('0')
    this.accounts.forEach(account => {
      const rightValue = Web3.utils.toBN(account.balance)
      value = value.add(rightValue)
    })
    this.totalBalance = value.toString()
  }

  uploadKeyStoreHandler(event: any) {
    const file = event.files[0]
    this.accountService.uploadKeyStoreFile(file, this.keystorePassPhrase).subscribe(res => {
      if (res.code === Code.ok) {
        this.msgService.addSuccess("导入keystore成功");
        this.keystoreImportVisible = false;
        this.load()
      } else {
        this.msgService.addError(`导入keystore失败: ${res.msg}`);
      }
    })
    console.log(file)
  }

  exportKeyStore() {
    this.msgService.confirm(`
        您将导出您的全部账户到keystore文件,
        导出格式为<font class="text-orange-500">zip压缩包</font>,
        keystore密钥为<font class="text-orange-500">您的钱包账户密码</font>, 确认导出?`,
      () => {
        this.loading = true;
        this.accountService.exportKeyStoreFile().subscribe(res => {
          this.loading = false;
          let dataType = res.type;
          let binaryData = [];
          binaryData.push(res);
          let downloadLink = document.createElement('a');
          downloadLink.href = window.URL.createObjectURL(new Blob(binaryData, {type: dataType}));
          downloadLink.setAttribute('download', "keystores.zip");
          document.body.appendChild(downloadLink);
          downloadLink.click();
          this.msgService.addSuccess("导出keystore成功");
        })
      })
  }

  importPrivateKey() {
    this.loading = true;
    this.accountService.uploadPrivateKey(this.privateKey).subscribe(res => {
      if (res.code === Code.ok) {
        this.msgService.addSuccess("导入私钥成功")
        this.privateKeyImportVisible = false;
        this.load()
      } else {
        this.msgService.addError(`导入私钥失败, error: ${res.msg}`)
      }
    })
  }

  newAccount() {
    this.msgService.confirm("系统将基于您的助记词为您新增账户, 请确认", () => {
      this.accountService.generateNewAccount().subscribe(res => {
        if (res.code === Code.ok) {
          this.msgService.addSuccess(`新增账户成功`)
          this.load()
        } else {
          this.msgService.addError(`新增账户失败, error: ${res.msg}`)
        }
      })
    })
  }
}
