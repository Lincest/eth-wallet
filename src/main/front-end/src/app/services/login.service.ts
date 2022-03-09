import {Injectable} from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {AUTH_URL, BASE_URL} from "../models/global";
import {map, Observable} from "rxjs";
import {Code, Resp} from "../models/resp";
import {ActivatedRouteSnapshot, CanActivate, Router, RouterStateSnapshot} from "@angular/router";

@Injectable({
  providedIn: 'root'
})
export class LoginService implements CanActivate {

  constructor(
    private http: HttpClient,
    private router: Router
  ) {
  }

  canActivate(route: ActivatedRouteSnapshot, state: RouterStateSnapshot) {
    return this.isLoggedIn().pipe(map(
      x => {
        if (!x) {
          this.router.navigate(['/login'], {queryParams: {returnUrl: state.url}}).then();
          return false;
        }
        return true;
      }
    ))
  }

  // check if logged in
  isLoggedIn(): Observable<boolean> {
    return this.http.get<Resp>(AUTH_URL + "/hello-world").pipe(
      map(x => {
        if (x.code === Code.ok) {
          localStorage.setItem("wallet-login", JSON.stringify(x.data))
        }
        return x.code === Code.ok
      })
    )
  }

  // 登出
  logOut(): Observable<boolean> {
    return this.http.post<Resp>(AUTH_URL + "/logout", {}).pipe(map(x => {
      if (x.code === Code.ok) {
        localStorage.removeItem("wallet-login");
        return true;
      }
      return false;
    }));
  }

  // 登录
  logIn(name: string, password: string): Observable<boolean> {
    return this.http.post<Resp>(BASE_URL + "/login", {name, password}).pipe(map(x => {
      if (x.code === Code.ok) {
        localStorage.setItem("wallet-login", JSON.stringify(x.data))
        return true;
      }
      return false;
    }))
  }
}
