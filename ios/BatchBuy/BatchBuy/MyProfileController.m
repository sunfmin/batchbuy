//
//  MyProfileController.m
//  BatchBuy
//
//  Created by Felix Sun on 10/15/13.
//  Copyright (c) 2013 HyperMusk. All rights reserved.
//

#import "MyProfileController.h"
#import "Profile.h"

@interface MyProfileController ()

@end

@implementation MyProfileController

- (id)initWithNibName:(NSString *)nibNameOrNil bundle:(NSBundle *)nibBundleOrNil
{
    self = [super initWithNibName:nibNameOrNil bundle:nibBundleOrNil];
    if (self) {
        // Custom initialization
    }
    return self;
}

- (IBAction)saveProfile:(id)sender {
    if ([[self.email.text stringByTrimmingCharactersInSet:[NSCharacterSet whitespaceAndNewlineCharacterSet]] compare: @""]== NSOrderedSame) {
        self.message.text=@"电子邮件不能为空.";
        return;
    }
    [Profile saveProfile:self.managedObjectContext name:self.name.text email:self.email.text];
    self.message.text= @"";
    [self.view endEditing:YES];
}

- (void)viewDidLoad
{
    Profile *p = [Profile loadProfile:self.managedObjectContext];
    if (p != nil) {
        self.name.text = p.name;
        self.email.text = p.email;
    } else {
        self.message.text=@"请设定你的姓名和电子邮件.";
    }
    [super viewDidLoad];
	// Do any additional setup after loading the view.
}

- (void)didReceiveMemoryWarning
{
    [super didReceiveMemoryWarning];
    // Dispose of any resources that can be recreated.
}

@end
