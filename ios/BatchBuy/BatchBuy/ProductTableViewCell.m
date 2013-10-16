//
//  ProductTableViewCell.m
//  BatchBuy
//
//  Created by Felix Sun on 10/14/13.
//  Copyright (c) 2013 HyperMusk. All rights reserved.
//

#import "ProductTableViewCell.h"
#import "Api.h"

@implementation ProductTableViewCell

- (IBAction)increaseCount:(id)sender {

    Service *s = [Service alloc];
    ServicePutOrderResults *r = [s PutOrder:[self.extraInfo objectForKey:@"date"] email:[self.extraInfo objectForKey:@"email"] productId:self.productId count:[NSNumber numberWithInt:1]];

    if (r.Err != nil) {
        NSLog(@"%@", r.Err);
        return;
    }

    self.orderCountValue = self.orderCountValue+1;
    self.orderCount.text = [NSString stringWithFormat:@"%ld", (long)self.orderCountValue];
}

- (IBAction)decreaseCount:(id)sender {
    
    Service *s = [Service alloc];
    NSError *err = [s RemoveOrder:[self.extraInfo objectForKey:@"date"] email:[self.extraInfo objectForKey:@"email"] productId:self.productId];

    if (err != nil) {
        NSLog(@"%@", err);
        return;
    }

    self.orderCountValue = self.orderCountValue-1;
    if (self.orderCountValue<0) {
        self.orderCountValue = 0;
    }
    self.orderCount.text = [NSString stringWithFormat:@"%ld", (long)self.orderCountValue];
}


- (id)initWithStyle:(UITableViewCellStyle)style reuseIdentifier:(NSString *)reuseIdentifier
{
    self = [super initWithStyle:style reuseIdentifier:reuseIdentifier];
    if (self) {
        // Initialization code
    }
    return self;
}

- (void)setSelected:(BOOL)selected animated:(BOOL)animated
{
    [super setSelected:selected animated:animated];

    // Configure the view for the selected state
}

@end
