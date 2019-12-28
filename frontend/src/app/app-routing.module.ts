import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { GroupGenerationComponent } from './components/group-generation/group-generation.component';


const routes: Routes = [
  {path: '', component: GroupGenerationComponent}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
