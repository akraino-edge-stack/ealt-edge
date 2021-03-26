import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { PcbimagesComponent } from './pcbimages.component';

describe('PcbimagesComponent', () => {
  let component: PcbimagesComponent;
  let fixture: ComponentFixture<PcbimagesComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ PcbimagesComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(PcbimagesComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
