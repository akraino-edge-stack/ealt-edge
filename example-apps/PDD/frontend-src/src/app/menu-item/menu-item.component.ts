import { Component, OnInit, ViewChild, Input } from '@angular/core';
import {Router} from '@angular/router';

@Component({
  selector: 'app-menu-item',
  templateUrl: './menu-item.component.html',
  styleUrls: ['./menu-item.component.scss']
})
export class MenuItemComponent implements OnInit {

  @Input() items: [];

  @ViewChild('childMenu', {static: true}) public childMenu;

  constructor(public router: Router) {
  }

  ngOnInit() {
  }

}
