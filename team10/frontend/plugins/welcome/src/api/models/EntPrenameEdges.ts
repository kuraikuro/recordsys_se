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
    EntPatientrecord,
    EntPatientrecordFromJSON,
    EntPatientrecordFromJSONTyped,
    EntPatientrecordToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntPrenameEdges
 */
export interface EntPrenameEdges {
    /**
     * Patientrecord holds the value of the patientrecord edge.
     * @type {Array<EntPatientrecord>}
     * @memberof EntPrenameEdges
     */
    patientrecord?: Array<EntPatientrecord>;
    /**
     * Prename2doctorinfo holds the value of the prename2doctorinfo edge.
     * @type {Array<EntDoctorinfo>}
     * @memberof EntPrenameEdges
     */
    prename2doctorinfo?: Array<EntDoctorinfo>;
}

export function EntPrenameEdgesFromJSON(json: any): EntPrenameEdges {
    return EntPrenameEdgesFromJSONTyped(json, false);
}

export function EntPrenameEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntPrenameEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'patientrecord': !exists(json, 'Patientrecord') ? undefined : ((json['Patientrecord'] as Array<any>).map(EntPatientrecordFromJSON)),
        'prename2doctorinfo': !exists(json, 'prename2doctorinfo') ? undefined : ((json['prename2doctorinfo'] as Array<any>).map(EntDoctorinfoFromJSON)),
    };
}

export function EntPrenameEdgesToJSON(value?: EntPrenameEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'patientrecord': value.patientrecord === undefined ? undefined : ((value.patientrecord as Array<any>).map(EntPatientrecordToJSON)),
        'prename2doctorinfo': value.prename2doctorinfo === undefined ? undefined : ((value.prename2doctorinfo as Array<any>).map(EntDoctorinfoToJSON)),
    };
}


