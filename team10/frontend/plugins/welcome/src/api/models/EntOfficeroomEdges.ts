/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntDoctorinfo,
    EntDoctorinfoFromJSON,
    EntDoctorinfoFromJSONTyped,
    EntDoctorinfoToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntOfficeroomEdges
 */
export interface EntOfficeroomEdges {
    /**
     * Officeroom2doctorinfo holds the value of the officeroom2doctorinfo edge.
     * @type {Array<EntDoctorinfo>}
     * @memberof EntOfficeroomEdges
     */
    officeroom2doctorinfo?: Array<EntDoctorinfo>;
}

export function EntOfficeroomEdgesFromJSON(json: any): EntOfficeroomEdges {
    return EntOfficeroomEdgesFromJSONTyped(json, false);
}

export function EntOfficeroomEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntOfficeroomEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'officeroom2doctorinfo': !exists(json, 'officeroom2doctorinfo') ? undefined : ((json['officeroom2doctorinfo'] as Array<any>).map(EntDoctorinfoFromJSON)),
    };
}

export function EntOfficeroomEdgesToJSON(value?: EntOfficeroomEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'officeroom2doctorinfo': value.officeroom2doctorinfo === undefined ? undefined : ((value.officeroom2doctorinfo as Array<any>).map(EntDoctorinfoToJSON)),
    };
}


