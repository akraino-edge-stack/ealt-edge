import { ComponentFixture, TestBed } from '@angular/core/testing';

import { BackuprestoreComponent } from './backuprestore.component';

describe('BackuprestoreComponent', () => {
  let component: BackuprestoreComponent;
  let fixture: ComponentFixture<BackuprestoreComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ BackuprestoreComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(BackuprestoreComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
