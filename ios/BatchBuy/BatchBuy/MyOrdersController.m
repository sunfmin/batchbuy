//
//  FirstViewController.m
//  BatchBuy
//
//  Created by Felix Sun on 10/14/13.
//  Copyright (c) 2013 HyperMusk. All rights reserved.
//

#import "MyOrdersController.h"
#import "ProductTableViewCell.h"
#import "Api.h"
#import "SDWebImage/UIImageView+WebCache.h"

@interface MyOrdersController ()

{
    
    NSArray *allProducts;
}
@end

@implementation MyOrdersController

- (NSInteger)numberOfSectionsInTableView:(UITableView *)tableView {
	// Number of sections is the number of regions.
	return 2;
}
- (NSInteger)tableView:(UITableView *)tableView numberOfRowsInSection:(NSInteger)section {
    if (section == 0) {
        return 1;
    } else {
        return allProducts.count;
    }
	return 0;
}

- (NSString *)tableView:(UITableView *)tableView titleForHeaderInSection:(NSInteger)section {
	// The header for the section is the region name -- get this from the region at the section index.
	if (section == 0) {
        return @"已下订单";
    } else {
        return @"其他点心";
    }
	return @"";
}

- (UITableViewCell *)tableView:(UITableView *)tableView cellForRowAtIndexPath:(NSIndexPath *)indexPath {
    
	static NSString *MyIdentifier = @"Cell";
    
	ProductTableViewCell *cell = [tableView dequeueReusableCellWithIdentifier:MyIdentifier];

    if (indexPath.section == 0) {
        
    } else {
        Product *p = [allProducts objectAtIndex:indexPath.row];
        cell.productTitle.text = p.Name;
        cell.orderCount.text = @"";
        [cell.imageView setImageWithURL:[NSURL URLWithString:p.PhotoLink]];
        
    }
	return cell;
}



- (void)viewDidLoad
{
    [super viewDidLoad];
    
    Service *s = [Service alloc];
    ServiceProductListOfDateResults *r = [s ProductListOfDate:@"2013-10-14"];

    if (r.Err != nil) {
        NSLog(@"%@", r.Err);
        return;
    }
    allProducts = r.Products;
}

- (void)didReceiveMemoryWarning
{
    [super didReceiveMemoryWarning];
    // Dispose of any resources that can be recreated.
}

@end
