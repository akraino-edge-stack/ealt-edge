import { Injectable } from '@angular/core';
// import { Observable } from 'rxjs/Observable';
import { Observable,throwError } from 'rxjs'
import { of } from 'rxjs';
// import {} from './../../assets/images/akraino-icon-1.png'
// import 'rxjs/add/Observable/of';
// import 'of;

@Injectable({
  providedIn: 'root'
})
export class ToolbarService {

  visible: Observable<boolean>;
  vis: boolean;

  constructor() { }

  hide() {
    this.visible = of(false);
    console.log("inside hide method");
    console.log(this.visible);
  }

  show() {
     this.visible = of(true);
     console.log(this.visible);
    //  console.log(this.visible.value);
    // this.vis = this.visible.value;
  }

  toggle() {
    this.visible =of(!this.visible);
  }
}
