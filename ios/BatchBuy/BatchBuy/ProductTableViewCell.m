//
//  ProductTableViewCell.m
//  BatchBuy
//
//  Created by Felix Sun on 10/14/13.
//  Copyright (c) 2013 HyperMusk. All rights reserved.
//

#import "ProductTableViewCell.h"

@implementation ProductTableViewCell

- (IBAction)increaseCount:(id)sender {
    self.orderCountValue = self.orderCountValue+1;
    self.orderCount.text = [NSString stringWithFormat:@"%ld", (long)self.orderCountValue];
}

- (IBAction)decreaseCount:(id)sender {
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
