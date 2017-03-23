import { RouterModule, Routes } from '@angular/router'

import { ProductListComponent } from './product-list/product-list.component'

const routes: Routes = [
  { path: 'product',

    children: [
      { path: 'list',
        component: ProductListComponent,
        data: { title: 'Lista produktów' }
      }
    ]
  }
];

export const ProductRoutes = RouterModule.forChild(routes);
