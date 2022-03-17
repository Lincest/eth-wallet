import {NgModule} from '@angular/core';
import {RouterModule, Routes} from '@angular/router';
import {LoginService} from "./services/login.service";
import {LoginComponent} from "./login/login.component";
import {HomeComponent} from "./home/home.component";
import {MnemonicComponent} from "./mnemonic/mnemonic.component";
import {NetworkComponent} from "./components/network/network.component";
import {AccountComponent} from "./components/account/account.component";
import {TransactionComponent} from "./components/transaction/transaction.component";

const routes: Routes = [
  {path: 'login', component: LoginComponent},
  {
    path: 'home', component: HomeComponent, canActivate: [LoginService], children: [
      {path: 'network', component: NetworkComponent},
      {path: 'account', component: AccountComponent},
      {path: 'transaction', component: TransactionComponent}
    ]
  },
  {path: 'mnemonic', component: MnemonicComponent, canActivate: [LoginService]},
  // default
  {path: '**', redirectTo: 'home'}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule {
}
