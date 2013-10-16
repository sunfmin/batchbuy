//
//  FirstViewController.h
//  BatchBuy
//
//  Created by Felix Sun on 10/14/13.
//  Copyright (c) 2013 HyperMusk. All rights reserved.
//

#import <UIKit/UIKit.h>
#import "Profile.h"

@interface MyOrdersController : UITableViewController

@property (nonatomic, strong) Profile *profile;
@property (weak, nonatomic) IBOutlet UINavigationItem *navItem;

@property NSDate *currentDate;
@property NSMutableDictionary *extraInfo;

@end
