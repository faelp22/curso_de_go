const getListMunicipiosByUF = (state) => (uf) => {
    if (state.lista_municipios_geral.hasOwnProperty(uf)) {
        return state.lista_municipios_geral[uf];
    }
    return [];
}

export default {
    getListMunicipiosByUF,
};
