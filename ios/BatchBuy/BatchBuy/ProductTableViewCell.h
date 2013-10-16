//
//  ProductTableViewCell.h
//  BatchBuy
//
//  Created by Felix Sun on 10/14/13.
//  Copyright (c) 2013 HyperMusk. All rights reserved.
//

#import <UIKit/UIKit.h>
#import "MyOrdersController.h"

@interface ProductTableViewCell : UITableViewCell
@property (weak, nonatomic) IBOutlet UIImageView *imageView;
@property (weak, nonatomic) IBOutlet UILabel *productTitle;
@property (weak, nonatomic) IBOutlet UILabel *orderCount;

@property NSInteger orderCountValue;
@property NSString *productId;

@property NSMutableDictionary *extraInfo;

@end
