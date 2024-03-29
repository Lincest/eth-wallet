import {Injectable} from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {AUTH_URL, BASE_URL} from "../models/global";
import {map, Observable} from "rxjs";
import {Code, Resp} from "../models/resp";
import {ActivatedRouteSnapshot, CanActivate, Router, RouterStateSnapshot} from "@angular/router";
import { CookieService } from 'ngx-cookie-service';

@Injectable({
  providedIn: 'root'
})
export class LoginService implements CanActivate {

  constructor(
    private http: HttpClient,
    private router: Router,
    private cookieService: CookieService,
  ) {
  }

  canActivate(route: ActivatedRouteSnapshot, state: RouterStateSnapshot) {
    return this.isLoggedIn().pipe(map(
      x => {
        if (!x) {
          this.router.navigate(['/login'], {queryParams: {returnUrl: state.url}}).then();
          return false;
        }
        this.checkMnemonic().subscribe(flag => {
          if (!flag) {
            this.router.navigate(['/mnemonic'], {queryParams: {returnUrl: state.url}}).then();
            return false;
          }
          return true;
        })
        return true;
      }
    ))
  }

  // check if logged in
  isLoggedIn(): Observable<boolean> {
    return this.http.get<Resp>(AUTH_URL + "/token").pipe(
      map(x => {
        if (x.code === Code.ok) {
          this.cookieService.set('CSRF-TOKEN', x.data.token);
        }
        return x.code === Code.ok
      })
    )
  }

  // 登出
  logOut(): Observable<Resp> {
    return this.http.post<Resp>(AUTH_URL + "/logout", {}).pipe(map(x => {
      if (x.code === Code.ok) {
        localStorage.removeItem("wallet-login");
        this.cookieService.delete('CSRF-TOKEN');
      }
      return x;
    }));
  }

  // 登录
  logIn(name: string, password: string): Observable<Resp> {
    return this.http.post<Resp>(BASE_URL + "/login", {name, password}).pipe(map(x => {
      if (x.code === Code.ok) {
        localStorage.setItem("wallet-login", JSON.stringify({name, password}));
      }
      return x;
    }))
  }

  // 注册
  register(name: string, password: string): Observable<Resp> {
    return this.http.post<Resp>(BASE_URL + "/register", {name, password})
  }

  // 注册助记词
  updateMnemonic(mnemonic: string): Observable<Resp> {
    return this.http.post<Resp>(AUTH_URL + "/mnemonic", {mnemonic})
  }

  // 查询助记词是否创建
  checkMnemonic(): Observable<boolean> {
    return this.http.get<Resp>(AUTH_URL + "/mnemonic")
      .pipe(
        map (res => {
          return res.code === Code.ok && res.data !== "";
        })
      )
  }
}
