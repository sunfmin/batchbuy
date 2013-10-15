//
//  Profile.h
//  BatchBuy
//
//  Created by Felix Sun on 10/15/13.
//  Copyright (c) 2013 HyperMusk. All rights reserved.
//

#import <Foundation/Foundation.h>
#import <CoreData/CoreData.h>


@interface Profile : NSManagedObject

@property (nonatomic, retain) NSString * email;
@property (nonatomic, retain) NSString * name;

+(Profile*) loadProfile:(NSManagedObjectContext*)context;
+(void) saveProfile:(NSManagedObjectContext*)context name:(NSString*)name email:(NSString*)email;
@end
