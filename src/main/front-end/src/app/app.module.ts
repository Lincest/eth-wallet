import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import {ButtonModule} from 'primeng/button';
import { HomeComponent } from './home/home.component';
import { HttpClientModule } from '@angular/common/http';
import { LoginComponent } from './login/login.component';
import {PasswordModule} from "primeng/password";
import {FormsModule} from "@angular/forms";
import {InputTextModule} from "primeng/inputtext";
import {BrowserAnimationsModule} from "@angular/platform-browser/animations";
import {MessageModule} from "primeng/message";
import {MessagesModule} from "primeng/messages";
import {ToastModule} from "primeng/toast";
import {RippleModule} from "primeng/ripple";

@NgModule({
  declarations: [
    AppComponent,
    HomeComponent,
    LoginComponent
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
    RippleModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
