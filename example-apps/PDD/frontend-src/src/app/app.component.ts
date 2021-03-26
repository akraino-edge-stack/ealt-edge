import { Component } from '@angular/core';
import { ToolbarService } from './toolbar/toolbar.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent {
  title = 'EliotUI';

  constructor(public toolbarService: ToolbarService) {}

  ngOnInit() {
    this.toolbarService.show();
  }

}
