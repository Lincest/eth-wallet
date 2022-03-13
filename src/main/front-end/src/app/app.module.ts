import {NgModule} from '@angular/core';
import {BrowserModule} from '@angular/platform-browser';

import {AppRoutingModule} from './app-routing.module';
import {AppComponent} from './app.component';
import {ButtonModule} from 'primeng/button';
import {HomeComponent} from './home/home.component';
import {HttpClientModule} from '@angular/common/http';
import {LoginComponent} from './login/login.component';
import {PasswordModule} from "primeng/password";
import {FormsModule} from "@angular/forms";
import {InputTextModule} from "primeng/inputtext";
import {BrowserAnimationsModule} from "@angular/platform-browser/animations";
import {MessageModule} from "primeng/message";
import {MessagesModule} from "primeng/messages";
import {ToastModule} from "primeng/toast";
import {RippleModule} from "primeng/ripple";
import {MnemonicComponent} from './mnemonic/mnemonic.component';
import { NewMnemonicComponent } from './mnemonic/new-mnemonic/new-mnemonic.component';
import { LoadMnemonicComponent } from './mnemonic/load-mnemonic/load-mnemonic.component';
import {ChipsModule} from "primeng/chips";
import {ConfirmDialogModule} from "primeng/confirmdialog";
import { TopbarComponent } from './home/topbar/topbar.component';
import { MenuComponent } from './home/menu/menu.component';
import { FooterComponent } from './home/footer/footer.component';
import {MenuItemComponent} from "./home/menu/menu-item/menu-item.component";
import {OverlayPanelModule} from 'primeng/overlaypanel';
import {Menu, MenuModule} from 'primeng/menu';
import {MenuItem} from 'primeng/api';
import {ChipModule} from "primeng/chip";
import {CardModule} from "primeng/card";

@NgModule({
  declarations: [
    AppComponent,
    HomeComponent,
    LoginComponent,
    MnemonicComponent,
    NewMnemonicComponent,
    LoadMnemonicComponent,
    TopbarComponent,
    MenuComponent,
    FooterComponent,
    MenuItemComponent
  ],
  imports: [
    BrowserModule,
    BrowserAnimationsModule,
    MessageModule,
    MessagesModule,
    ToastModule,
    HttpClientModule,
    AppRoutingModule,
    PasswordModule,
    FormsModule,
    InputTextModule,
    ButtonModule,
    RippleModule,
    ChipsModule,
    ConfirmDialogModule,
    OverlayPanelModule,
    MenuModule,
    ChipModule,
    CardModule
  ],
  providers: [],
  bootstrap: [AppComponent],
})
export class AppModule {}
