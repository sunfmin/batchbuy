//
//  Profile.m
//  BatchBuy
//
//  Created by Felix Sun on 10/15/13.
//  Copyright (c) 2013 HyperMusk. All rights reserved.
//

#import "Profile.h"


@implementation Profile

@dynamic email;
@dynamic name;

+(Profile*) loadProfile:(NSManagedObjectContext*)context{
    NSFetchRequest *fetchRequest = [[NSFetchRequest alloc] init];
    NSEntityDescription *entity = [NSEntityDescription entityForName:@"Profile" inManagedObjectContext:context];
    [fetchRequest setEntity:entity];
    NSError *error;
    NSArray *rs = [context executeFetchRequest:fetchRequest error:&error];
    
    if (rs.count == 0) {
        return nil;
    }
    
    return [rs objectAtIndex:0];
}

+(void) saveProfile:(NSManagedObjectContext*)context name:(NSString*)name email:(NSString*)email{
    
    Profile *profile = [Profile loadProfile:context];
    BOOL insert = NO;
    if (profile == nil) {
        profile = (Profile *)[NSEntityDescription insertNewObjectForEntityForName:@"Profile" inManagedObjectContext:context];
        insert = YES;
    }
    
    [profile setEmail:email];
    [profile setName:name];
    
    if (insert) {
        [context insertObject:profile];
    } else {
        [context refreshObject:profile mergeChanges:YES];
    }

    NSError *error;
    [context save:&error];
}

static NSDateFormatter * _dateFormatter;

+ (NSDateFormatter *) dateFormatter {
	if(!_dateFormatter) {
		_dateFormatter = [[NSDateFormatter alloc] init];
		[_dateFormatter setDateFormat:@"yyyy-MM-dd"];
	}
	return _dateFormatter;
}

@end
