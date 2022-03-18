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
import {ChipsModule} from "primeng/chips";
import {ConfirmDialogModule} from "primeng/confirmdialog";
import {TopbarComponent} from './home/topbar/topbar.component';
import {MenuComponent} from './home/menu/menu.component';
import {FooterComponent} from './home/footer/footer.component';
import {MenuItemComponent} from "./home/menu/menu-item/menu-item.component";
import {OverlayPanelModule} from 'primeng/overlaypanel';
import {MenuModule} from 'primeng/menu';
import {ChipModule} from "primeng/chip";
import {CardModule} from "primeng/card";
import {NetworkComponent} from './components/network/network.component';
import {AccordionModule} from "primeng/accordion";
import {DividerModule} from "primeng/divider";
import {DialogModule} from "primeng/dialog";
import { AccountComponent } from './components/account/account.component';
import {TableModule} from "primeng/table";
import {ToolbarModule} from "primeng/toolbar";
import {FileUpload, FileUploadModule} from "primeng/fileupload";
import {ProgressBarModule} from "primeng/progressbar";
import { EthWeiPipe } from './pipes/eth-wei.pipe';
import { TransactionComponent } from './components/transaction/transaction.component';
import {AutoCompleteModule} from "primeng/autocomplete";
import {InputNumberModule} from "primeng/inputnumber";
import { TransactionItemComponent } from './components/transaction-search/transaction-item/transaction-item.component';
import { PageNotFoundComponent } from './page-not-found/page-not-found.component';
import { HomeDefaultComponent } from './components/home-default/home-default.component';
import { TransactionSearchComponent } from './components/transaction-search/transaction-search.component';

@NgModule({
  declarations: [
    AppComponent,
    HomeComponent,
    LoginComponent,
    MnemonicComponent,
    TopbarComponent,
    MenuComponent,
    FooterComponent,
    MenuItemComponent,
    NetworkComponent,
    AccountComponent,
    EthWeiPipe,
    TransactionComponent,
    TransactionItemComponent,
    PageNotFoundComponent,
    HomeDefaultComponent,
    TransactionSearchComponent
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
        CardModule,
        AccordionModule,
        DividerModule,
        InputTextModule,
        DialogModule,
        TableModule,
        ToolbarModule,
        FileUploadModule,
        ProgressBarModule,
        AutoCompleteModule,
        InputNumberModule
    ],
  providers: [],
  bootstrap: [AppComponent],
})
export class AppModule {}
