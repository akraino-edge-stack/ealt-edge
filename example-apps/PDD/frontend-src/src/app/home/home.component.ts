import { Component, OnInit } from '@angular/core';
import { ToolbarService } from './../toolbar/toolbar.service';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.scss']
})
export class HomeComponent implements OnInit {
  // images = [944, 1011, 984].map((n) => `https://picsum.photos/id/${n}/900/500`);
  images = ["./../../assets/images/iotlatest1.png","./../../assets/images/iotlatest1.png","./../../assets/images/iotlatest1.png"];
  public currentUser;
  
  constructor(public toolbarService: ToolbarService) {
    this.currentUser = localStorage.getItem('currentUser')? JSON.parse(localStorage.getItem('currentUser')) : '';
  }

  ngOnInit() {
    // this.toolbarService.show();
    this.toolbarService.show();
  }

}
