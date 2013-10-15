//
//  MyProfileController.h
//  BatchBuy
//
//  Created by Felix Sun on 10/15/13.
//  Copyright (c) 2013 HyperMusk. All rights reserved.
//

#import <UIKit/UIKit.h>
#import <CoreData/CoreData.h>

@interface MyProfileController : UIViewController
@property (weak, nonatomic) IBOutlet UITextField *name;
@property (weak, nonatomic) IBOutlet UITextField *email;
@property (weak, nonatomic) IBOutlet UILabel *message;
@property (nonatomic, strong) NSManagedObjectContext *managedObjectContext;


@end
