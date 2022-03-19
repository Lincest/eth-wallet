import {NgModule} from '@angular/core';
import {RouterModule, Routes} from '@angular/router';
import {LoginService} from "./services/login.service";
import {LoginComponent} from "./login/login.component";
import {HomeComponent} from "./home/home.component";
import {MnemonicComponent} from "./mnemonic/mnemonic.component";
import {NetworkComponent} from "./components/network/network.component";
import {AccountComponent} from "./components/account/account.component";
import {TransactionComponent} from "./components/transaction/transaction.component";
import {TransactionItemComponent} from "./components/transaction-search/transaction-item/transaction-item.component";
import {PageNotFoundComponent} from "./page-not-found/page-not-found.component";
import {HomeDefaultComponent} from "./components/home-default/home-default.component";
import {TransactionSearchComponent} from "./components/transaction-search/transaction-search.component";
import {TransactionHistoryComponent} from "./components/transaction-history/transaction-history.component";

const routes: Routes = [
  {path: 'login', component: LoginComponent},
  {
    path: 'home', component: HomeComponent, canActivate: [LoginService], children: [
      {path: '', component: HomeDefaultComponent},
      {path: 'network', component: NetworkComponent},
      {path: 'account', component: AccountComponent},
      {path: 'transaction', component: TransactionComponent},
      {
        path: 'transaction-search', component: TransactionSearchComponent, children: [
          {path: ':hash', component: TransactionItemComponent},
        ]
      },
      {path: 'transaction-history', component: TransactionHistoryComponent}
    ]
  },
  {path: 'mnemonic', component: MnemonicComponent, canActivate: [LoginService]},
  // default
  {path: '', redirectTo: 'home', pathMatch: 'full'},
  {path: '404', component: PageNotFoundComponent},
  {path: '**', component: PageNotFoundComponent}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule {
}
