import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { PcboutputimagesComponent } from './pcboutputimages.component';

describe('PcboutputimagesComponent', () => {
  let component: PcboutputimagesComponent;
  let fixture: ComponentFixture<PcboutputimagesComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ PcboutputimagesComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(PcboutputimagesComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
