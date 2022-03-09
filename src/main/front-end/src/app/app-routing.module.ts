import {NgModule} from '@angular/core';
import {RouterModule, Routes} from '@angular/router';
import {AppComponent} from "./app.component";
import {LoginService} from "./services/login.service";
import {LoginComponent} from "./login/login.component";
import {HomeComponent} from "./home/home.component";
import {MnemonicComponent} from "./mnemonic/mnemonic.component";
import {NewMnemonicComponent} from "./mnemonic/new-mnemonic/new-mnemonic.component";
import {LoadMnemonicComponent} from "./mnemonic/load-mnemonic/load-mnemonic.component";

const routes: Routes = [
  {path: 'login', component: LoginComponent},
  {path: 'home', component: HomeComponent, canActivate: [LoginService]},
  {
    path: 'mnemonic', component: MnemonicComponent, canActivate: [LoginService], children: [
      {path: 'new', component: NewMnemonicComponent},
      {path: 'load', component: LoadMnemonicComponent},
    ]
  },
  // default
  {path: '**', redirectTo: 'home'}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule {
}
