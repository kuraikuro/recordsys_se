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
    EntTreatmentEdges,
    EntTreatmentEdgesFromJSON,
    EntTreatmentEdgesFromJSONTyped,
    EntTreatmentEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntTreatment
 */
export interface EntTreatment {
    /**
     * Datetreat holds the value of the "Datetreat" field.
     * @type {string}
     * @memberof EntTreatment
     */
    datetreat?: string;
    /**
     * Treatment holds the value of the "Treatment" field.
     * @type {string}
     * @memberof EntTreatment
     */
    treatment?: string;
    /**
     * 
     * @type {EntTreatmentEdges}
     * @memberof EntTreatment
     */
    edges?: EntTreatmentEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntTreatment
     */
    id?: number;
}

export function EntTreatmentFromJSON(json: any): EntTreatment {
    return EntTreatmentFromJSONTyped(json, false);
}

export function EntTreatmentFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntTreatment {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'datetreat': !exists(json, 'Datetreat') ? undefined : json['Datetreat'],
        'treatment': !exists(json, 'Treatment') ? undefined : json['Treatment'],
        'edges': !exists(json, 'edges') ? undefined : EntTreatmentEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
    };
}

export function EntTreatmentToJSON(value?: EntTreatment | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'Datetreat': value.datetreat,
        'Treatment': value.treatment,
        'edges': EntTreatmentEdgesToJSON(value.edges),
        'id': value.id,
    };
}


