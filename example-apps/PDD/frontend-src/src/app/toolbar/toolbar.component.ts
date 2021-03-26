import { Component, OnInit } from '@angular/core';
import { ToolbarService } from './toolbar.service';

@Component({
  selector: 'app-toolbar',
  templateUrl: './toolbar.component.html',
  styleUrls: ['./toolbar.component.scss']
})
export class ToolbarComponent implements OnInit {
  toolbarmenu = [];
  menuicon: string;
  constructor(public toolbarService: ToolbarService) { }

  ngOnInit() {
    this.menuicon = "menu";
    this.init();
  }

  init(){
    this.toolbarmenu = [      
      {
        displayName: 'EALT-EDGE',
        route: '/home'
      },
      // {
      //   displayName: 'PCB USECASE',
      // },
      {
        displayName: 'USECASES',
        children: [
          {
            displayName: 'PCB Defect Detection',
            route: '/pcb'
          }
        ]
      },
      {
        displayName: 'Akraino EALT Edge Wiki',
        route: '/akrainowiki'
      },
    ]
  }

}
